package metadata

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/bububa/opentaobao/metadata/util"
)

// ApiParamType 淘宝API文档字段类型
type ApiParamType = string

const (
	STRING_PARAM_TYPE     ApiParamType = "String"
	NUMBER_PARAM_TYPE     ApiParamType = "Number"
	DATE_PARAM_TYPE       ApiParamType = "Date"
	FIELD_LIST_PARAM_TYPE ApiParamType = "Field List"
	BOOLEAN_PARAM_TYPE    ApiParamType = "Boolean"
	PRICE_PARAM_TYPE      ApiParamType = "Price"
	BYTE_PARAM_TYPE       ApiParamType = "byte"
	BYTES_PARAM_TYPE      ApiParamType = "byte[]"
	JSON_PARAM_TYPE       ApiParamType = "Json"
	RESULT_PARAM_TYPE     ApiParamType = "Result"
	UNKNOWN_PARAM_TYPE    ApiParamType = ""
)

// ApiDocResponse 下载淘宝API文档返回结果
type ApiDocResponse struct {
	Success bool   `json:"success,omitempty"` // 是否成功
	Code    string `json:"code,omitempty"`    // 错误代码
	ErrMsg  string `json:"errMsg,omitempty"`  // 错误信息
	Data    ApiDoc `json:"data,omitempty"`    // 成功数据
}

// IsError 判断是否为error
func (res ApiDocResponse) IsError() bool {
	return !res.Success
}

// Error implement Error interface
func (res ApiDocResponse) Error() string {
	return res.ErrMsg
}

// ApiDoc 淘宝API文档结构体
type ApiDoc struct {
	Id             int64      `json:"id,omitempty"`             // API ID
	Name           string     `json:"name,omitempty"`           // API 方法名
	ChineseName    string     `json:"apiChineseName,omitempty"` // API 中文名
	Description    string     `json:"description,omitempty"`    // API 描述
	RequestParams  []ApiParam `json:"requestParams,omitempty"`  // 请求参数
	ResponseParams []ApiParam `json:"responseParams,omitempty"` // 返回参数
}

// Filename 获取淘宝API文档保存文件名
func (d ApiDoc) Filename() string {
	return strings.TrimSpace(strings.ReplaceAll(d.Name, ".", "_"))
}

// Title 获取淘宝API文档对应golang API名
func (d ApiDoc) Title() string {
	name := strings.Title(strings.ReplaceAll(d.Name, "-", "_"))
	return strings.TrimSpace(strings.ReplaceAll(name, ".", ""))
}

// ApiTpl 淘宝API文档转golang API 模版结构体
func (d ApiDoc) ApiTpl() ApiTpl {
	snakeName := strings.ReplaceAll(d.Name, ".", "_")
	tpl := ApiTpl{
		Name:           d.Title(),
		ApiName:        d.Name,
		SnakeName:      snakeName,
		ResponseKey:    fmt.Sprintf("%s_response", strings.TrimPrefix(snakeName, "taobao_")),
		ChineseName:    d.ChineseName,
		Desc:           clearDesc(d.Description, false),
		RequestParams:  make([]TplParam, 0, len(d.RequestParams)),
		ResponseParams: make([]TplParam, 0, len(d.ResponseParams)),
	}
	reqParamMp := make(map[string]struct{}, len(d.RequestParams))
	for _, p := range d.RequestParams {
		param := p.TplParam(tpl.Name)
		if _, found := reqParamMp[param.Name]; found {
			continue
		}
		reqParamMp[param.Name] = struct{}{}
		tpl.RequestParams = append(tpl.RequestParams, param)
		if !tpl.IsMultipart && param.IsMultipart() {
			tpl.IsMultipart = true
		}
	}
	respParamMp := make(map[string]struct{}, len(d.ResponseParams))
	for _, p := range d.ResponseParams {
		param := p.TplParam(tpl.Name)
		if _, found := respParamMp[param.Name]; found {
			continue
		}
		respParamMp[param.Name] = struct{}{}
		if param.Name == "RequestId" {
			tpl.HasRequestId = true
		}
		tpl.ResponseParams = append(tpl.ResponseParams, param)
	}
	return tpl
}

// ApiParam is 淘宝API文档字段结构体
type ApiParam struct {
	Name        string       `json:"name,omitempty"`        // 参数名
	Type        ApiParamType `json:"type,omitempty"`        // 参数类型
	Description string       `json:"description,omitempty"` // 描述
	Required    bool         `json:"required,omitempty"`    // 是否必须
	MaxLength   json.Number  `json:"maxLength,omitempty"`   // 最大长度
	MaxListSize json.Number  `json:"maxListSize,omitempty"` // 列表最大数量
	MaxValue    json.Number  `json:"maxValue,omitempty"`    // 最大值
	MinValue    json.Number  `json:"minValue,omitempty"`    // 最小值
	SubParams   []ApiParam   `json:"subParams,omitempty"`   // 对象包含参数
	UsePointer  bool         `json:"use_pointer,omitempty"` // 生成SDK是否使用指针
}

// TplParam 转化为API模版参数结构体
func (p ApiParam) TplParam(apiName string) TplParam {
	name := strings.ReplaceAll(strings.Title(strings.ReplaceAll(p.Name, "_", ".")), ".", "")
	param := TplParam{
		Name:     name,
		Label:    strings.ToLower(string(name[0])) + name[1:],
		ParamKey: p.Name,
		Desc:     clearDesc(p.Description, true),
		Required: p.Required,
	}
	paramType := strings.TrimSpace(strings.TrimSuffix(p.Type, "[]"))
	paramType = convertConflictModelName(apiName, paramType)
	switch paramType {
	case STRING_PARAM_TYPE, UNKNOWN_PARAM_TYPE:
		param.Type = "string"
	case NUMBER_PARAM_TYPE:
		param.Type = "int64"
	case DATE_PARAM_TYPE:
		param.Type = "string"
	case BOOLEAN_PARAM_TYPE:
		param.Type = "bool"
	case PRICE_PARAM_TYPE:
		param.Type = "float64"
	case FIELD_LIST_PARAM_TYPE:
		param.Type = "[]string"
	case BYTES_PARAM_TYPE, BYTE_PARAM_TYPE:
		param.Type = "*model.File"
	case JSON_PARAM_TYPE:
		param.Type = "string"
	default:
		paramType = strings.Title(paramType)
		replaceMp := map[string]string{
			"DTO": "Dto",
			"DTo": "Dto",
			"DO":  "Do",
			"RS":  "Rs",
		}
		for fromSuffix, toSuffix := range replaceMp {
			if strings.HasSuffix(paramType, fromSuffix) {
				paramType = strings.TrimSuffix(paramType, fromSuffix) + toSuffix
				break
			}
		}
		param.ObjType = paramType
		param.SnakeType = util.SnakeCase(paramType)
		if strings.HasSuffix(p.Type, "[]") {
			param.Type = fmt.Sprintf("[]%s", paramType)
			param.IsList = true
		} else {
			param.Type = fmt.Sprintf("*%s", paramType)
		}
		if len(p.SubParams) == 0 {
			log.Printf("[WRN] name:%s, type:%s\n", p.Name, paramType)
		} else {
			keys := make(map[string]struct{}, len(p.SubParams))
			for _, par := range p.SubParams {
				subPar := par.TplParam(apiName)
				if _, found := keys[subPar.Name]; found {
					continue
				}
				keys[subPar.Name] = struct{}{}
				param.Obj = append(param.Obj, subPar)
			}
		}
		param.IsObject = true
	}
	if p.Type != BYTES_PARAM_TYPE && strings.HasSuffix(p.Type, "[]") && !strings.HasPrefix(param.Type, "[]") {
		param.ObjType = paramType
		param.SnakeType = param.Type
		param.Type = fmt.Sprintf("[]%s", param.Type)
		param.IsList = true
	}
	if p.UsePointer && !strings.HasPrefix(param.Name, "*") {
		param.Type = fmt.Sprintf("*%s", param.Type)
	}
	return param
}

// convertConflictModelName 理冲突结构体名(conflict_models.json内设置)
func convertConflictModelName(apiName string, name string) string {
	for _, n := range ConflictModels {
		if name == n {
			return fmt.Sprintf("%s%s", apiName, n)
		}
	}
	return name
}

// clearDesc 处理描述字段以匹配golang结构体注释
func clearDesc(desc string, inline bool) string {
	desc = strings.ReplaceAll(desc, "/*", "")
	desc = strings.ReplaceAll(desc, "*/", "")
	if inline {
		desc = strings.ReplaceAll(desc, "\n", "")
	} else {
		lines := strings.Split(desc, "\n")
		var arr []string
		for _, line := range lines {
			arr = append(arr, fmt.Sprintf("// %s", line))
		}
		desc = strings.Join(arr, "\n")
	}
	return desc
}
