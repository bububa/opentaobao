package metadata

import (
	"strings"
)

// ApiTpl SDK API模版结构体
type ApiTpl struct {
	Pkg            string     // 包名
	ApiName        string     // Api方法名
	Name           string     // Api结构体名
	ResponseKey    string     // ***_response
	SnakeName      string     // snake 化方法名
	ChineseName    string     // 中文名
	Desc           string     // 描述
	IsMultipart    bool       // 是否使用multipart/form
	RequestParams  []TplParam // 请求参数
	ResponseParams []TplParam // 返回参数
	HasRequestId   bool       // 返回参数中是否已包含RequestId
}

// TplParam SDK API模版字段参数结构体
type TplParam struct {
	Name      string     // 结构体成员名
	Label     string     // 第一个字母小写的
	ParamKey  string     // api 参数的key(json/xml)
	Type      string     // 参数类型
	ObjType   string     // 对象类型的对象名
	SnakeType string     // 对象类型名snake化
	Desc      string     // 描述
	IsObject  bool       // 是否是对象
	IsList    bool       // 是否是数组
	Required  bool       // 是否必须
	Obj       []TplParam // 对象类型结构体参数
}

// IsMultipart 判断SDK API是否需要使用multipart/form post
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

// TplModel SDK model模版结构体
type TplModel struct {
	Pkg         string     // 包名
	Name        string     // struct 名
	Params      []TplParam // struct 成员
	ImportModel bool       // 是否import github.com/bububa/opentaobao/model
}

// NeedImportModel 判断SDK model模版是否需要引用github.com/bububa/opentaobao/model包
func (t TplModel) NeedImportModel() bool {
	for _, p := range t.Params {
		if strings.HasSuffix(p.Type, "model.File") {
			return true
		}
	}
	return false
}

// ExtractModels 提取包内包含的结构体
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

// MergeModels 合并包内命名相同结构体
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
