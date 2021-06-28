package metadata

import (
	"strings"
)

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
}

type TplModel struct {
	Pkg    string
	Name   string
	Params []TplParam
}

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

func (p TplParam) IsMultipart() bool {
	if len(p.Obj) == 0 {
		return p.Type == "[]byte"
	}
	for _, obj := range p.Obj {
		if obj.IsMultipart() {
			return true
		}
	}
	return false
}

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

func MergeModels(models []TplModel) []TplModel {
	mp := make(map[string]*TplModel, len(models))
	for _, model := range models {
		model := model
		if m, found := mp[model.Name]; found {
			keys := make(map[string]struct{}, len(m.Params))
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
