package service

import (
	"errors"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"postman-to-md/app/constant"
	"postman-to-md/app/model"
)

var PmReader = &pmReaderService{}

type pmReaderService struct {
}

func (s *pmReaderService) Run(param *model.RunParam) (err error) {
	contentObj := &model.Collection{}
	err = s.ReadJsonFile(param, contentObj)
	if err != nil {
		return err
	}
	err = s.DataHandler(param, contentObj)
	if err != nil {
		return err
	}
	return nil
}

func (s *pmReaderService) ReadJsonFile(param *model.RunParam, contentObj *model.Collection) (err error) {
	if !gfile.IsFile(param.InputPath) {
		return errors.New("failed to open the file. Please check whether the file path is correct")
	}
	fileContent := gfile.GetContents(param.InputPath)
	if fileContent == "" {
		return errors.New("the content of file is empty")
	}
	contentJson, err := gjson.Decode(fileContent)
	if err != nil {
		return err
	}
	if contentObj == nil {
		contentObj = &model.Collection{}
	}
	err = gconv.Scan(contentJson, contentObj)
	if err != nil {
		return err
	}
	return nil
}

func (s *pmReaderService) DataHandler(param *model.RunParam, contentObj *model.Collection) (err error) {
	rootPath := param.OutputPath
	if rootPath == "" && constant.Env == "dev" {
		rootPath += "temp"
	}
	if rootPath != "" {
		if !gfile.IsDir(rootPath) {
			return errors.New(rootPath + " is not dir")
		}
		if gstr.PosRRune(rootPath, "/") != gstr.LenRune(rootPath)-1 {
			rootPath += "/"
		}
	}
	rootPath += "output_" + gtime.Now().Format("YmdHis")
	err = gfile.Mkdir(rootPath)
	if err != nil {
		return err
	}
	// 首先生成一个主readme.md文件
	readmeContent := MdMaker.Title(contentObj.Info.Name, 1)
	readmeContent += MdMaker.Text(contentObj.Info.Description)
	err = gfile.PutContents(rootPath+"/readme.md", readmeContent)
	if err != nil {
		return err
	}
	err = s.DataItemsHandler(rootPath, contentObj.Item, 2, readmeContent)
	if err != nil {
		return err
	}

	return nil
}

func (s *pmReaderService) DataItemsHandler(path string, items []*model.Item, level int, initContent ...string) (err error) {
	content := ""
	if len(initContent) > 0 {
		content += initContent[0]
	}
	for _, item := range items {
		if len(item.Item) > 0 {
			childDirPath := path + "/" + item.Name
			err = gfile.Mkdir(childDirPath)
			if err != nil {
				return err
			}
			fileInitContent := ""
			fileInitContent += MdMaker.Title(item.Name, level)
			fileInitContent += MdMaker.Text(item.Description)
			err = s.DataItemsHandler(childDirPath, item.Item, level+1, fileInitContent)
		} else {
			content += s.DataLeafHandler(item, level)
			content += "\n"
		}
	}
	err = gfile.PutContents(path+"/readme.md", content)
	if err != nil {
		return err
	}
	return err
}

func (s *pmReaderService) DataLeafHandler(item *model.Item, level int) (re string) {
	// title & description
	re += MdMaker.Title(item.Name, level)
	re += MdMaker.Text(item.Description)

	// request method
	color := "#0FB348"
	switch gstr.ToUpper(item.Request.Method) {
	case "GET":
		color = "#0FB348"
	case "POST":
		color = "#FFCC65"
	default:
		color = "#0FB348"
	}
	re += MdMaker.Bold(item.Request.Method) + "    " + MdMaker.Color(MdMaker.Bold(item.Name), color) + "\n"

	// request url
	re += MdMaker.Code(item.Request.Url.Raw)
	re += MdMaker.Text(item.Request.Description)

	// request header
	re += MdMaker.Title("Request Headers", 6)
	headerTitleSlice := []string{
		"Key",
		"Value",
	}
	headerDataSlice := make([][]string, 0)
	for _, headerItem := range item.Request.Header {
		itemKey := MdMaker.Bold(headerItem.Key)
		itemValue := headerItem.Value
		if headerItem.Description != "" {
			itemValue += "<br>" + headerItem.Description
		}
		tempData := make([]string, 0)
		tempData = append(tempData, itemKey, itemValue)
		headerDataSlice = append(headerDataSlice, tempData)
	}
	re += MdMaker.Table(headerTitleSlice, headerDataSlice)

	// query params
	if item.Request.Url.Query != nil && len(item.Request.Url.Query) > 0 {
		re += MdMaker.Title("Query Params", 6)
		queryParamsTitleSlice := []string{
			"Key",
			"Value",
		}
		queryParamsDataSlice := make([][]string, 0)
		for _, queryParamsItem := range item.Request.Url.Query {
			itemKey := MdMaker.Bold(queryParamsItem.Key)
			itemValue := queryParamsItem.Value
			if queryParamsItem.Description != "" {
				itemValue += "<br>" + queryParamsItem.Description
			}
			tempData := make([]string, 0)
			tempData = append(tempData, itemKey, itemValue)
			queryParamsDataSlice = append(queryParamsDataSlice, tempData)
		}
		re += MdMaker.Table(queryParamsTitleSlice, queryParamsDataSlice)
	}

	// todo Example

	if item.Request.Body != nil && item.Request.Body.Raw != "" {
		re += MdMaker.Title("Body", 6)
		re += MdMaker.Code(item.Request.Body.Raw)
	}

	if item.Response != nil && len(item.Response) > 0 {
		for _, responseItem := range item.Response {
			if responseItem.Header != nil {
				re += MdMaker.Title("Response", 6)
				re += MdMaker.Title("ResponseHeader", 6)
				responseHeaderTitleSlice := []string{
					"Key",
					"Value",
				}
				responseHeaderDataSlice := make([][]string, 0)
				for _, headerItem := range responseItem.Header {
					tempData := make([]string, 0)
					tempData = append(tempData, headerItem.Key, headerItem.Value)
					responseHeaderDataSlice = append(responseHeaderDataSlice, tempData)
				}
				re += MdMaker.Table(responseHeaderTitleSlice, responseHeaderDataSlice)
			}

			if responseItem.Body != "" {
				re += MdMaker.Title("ResponseBody", 6)
				re += MdMaker.Code(responseItem.Body)
			}
		}
	}

	return re
}
