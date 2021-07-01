package metadata

import (
	"strings"
)

// SDK API模版结构体
type ApiTpl struct {
	Pkg            string
	ApiName        string
	Name           string
	ResponseKey    string
	SnakeName      string
	ChineseName    string
	Desc           string
	IsMultipart    bool
	RequestParams  []TplParam
	ResponseParams []TplParam
	HasRequestId   bool
}

// SDK API模版字段参数结构体
type TplParam struct {
	Name      string
	Label     string
	ParamKey  string
	Type      string
	ObjType   string
	SnakeType string
	Desc      string
	IsObject  bool
	IsList    bool
	Required  bool
	Obj       []TplParam
}

// SDK API是否需要使用form/multipart post
func (p TplParam) IsMultipart() bool {
	if strings.HasSuffix(p.Type, "model.File") {
		return true
	}
	for _, obj := range p.Obj {
		if obj.IsMultipart() {
			return true
		}
	}
	return false
}

// SDK model模版结构体
type TplModel struct {
	Pkg         string
	Name        string
	Params      []TplParam
	ImportModel bool
}

// SDK model模版是否需要引用github.com/bububa/opentaobao/model包
func (t TplModel) NeedImportModel() bool {
	for _, p := range t.Params {
		if strings.HasSuffix(p.Type, "model.File") {
			return true
		}
	}
	return false
}

// SDK 提取包内包含的结构体
func ExtractModels(params []TplParam) []TplModel {
	var models []TplModel
	for _, p := range params {
		if p.IsObject {
			objType := strings.TrimPrefix(strings.TrimPrefix(p.Type, "[]"), "*")
			model := TplModel{
				Name:   objType,
				Params: p.Obj,
			}
			models = append(models, model)
		}
		if len(p.Obj) > 0 {
			subModules := ExtractModels(p.Obj)
			models = append(models, subModules...)
		}
	}
	return models
}

// SDK 合并包内命名相同结构体
func MergeModels(models []TplModel) []TplModel {
	mp := make(map[string]*TplModel, len(models))
	for _, model := range models {
		model := model
		if m, found := mp[model.Name]; found {
			keys := make(map[string]struct{})
			for _, param := range m.Params {
				keys[param.Name] = struct{}{}
			}
			for _, param := range model.Params {
				if _, found := keys[param.Name]; !found {
					keys[param.Name] = struct{}{}
					m.Params = append(m.Params, param)
				}
			}
			continue
		}
		mp[model.Name] = &model
	}
	var ret []TplModel
	for _, model := range mp {
		ret = append(ret, *model)
	}
	return ret
}
