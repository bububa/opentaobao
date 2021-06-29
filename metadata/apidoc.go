package metadata

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/bububa/opentaobao/metadata/util"
)

type ApiParamType = string

const (
	STRING_PARAM_TYPE     ApiParamType = "String"
	NUMBER_PARAM_TYPE     ApiParamType = "Number"
	DATE_PARAM_TYPE       ApiParamType = "Date"
	FIELD_LIST_PARAM_TYPE ApiParamType = "Field List"
	BOOLEAN_PARAM_TYPE    ApiParamType = "Boolean"
	PRICE_PARAM_TYPE      ApiParamType = "Price"
	BYTES_PARAM_TYPE      ApiParamType = "byte"
	JSON_PARAM_TYPE       ApiParamType = "Json"
	RESULT_PARAM_TYPE     ApiParamType = "Result"
	UNKNOWN_PARAM_TYPE    ApiParamType = ""
)

type ApiDocResponse struct {
	Success bool   `json:"success,omitempty"`
	Code    string `json:"code,omitempty"`
	ErrMsg  string `json:"errMsg,omitempty"`
	Data    ApiDoc `json:"data,omitempty"`
}

func (res ApiDocResponse) IsError() bool {
	return !res.Success
}

func (res ApiDocResponse) Error() string {
	return res.ErrMsg
}

type ApiDoc struct {
	Id             int64      `json:"id,omitempty"`
	Name           string     `json:"name,omitempty"`
	ChineseName    string     `json:"apiChineseName,omitempty"`
	Description    string     `json:"description,omitempty"`
	RequestParams  []ApiParam `json:"requestParams,omitempty"`
	ResponseParams []ApiParam `json:"responseParams,omitempty"`
}

func (d ApiDoc) Filename() string {
	return strings.TrimSpace(strings.ReplaceAll(d.Name, ".", "_"))
}

func (d ApiDoc) Title() string {
	name := strings.Title(strings.ReplaceAll(d.Name, "-", "_"))
	return strings.TrimSpace(strings.ReplaceAll(name, ".", ""))
}

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

type ApiParam struct {
	Name        string       `json:"name,omitempty"`
	Type        ApiParamType `json:"type,omitempty"`
	Description string       `json:"description,omitempty"`
	Required    bool         `json:"required,omitempty"`
	MaxLength   json.Number  `json:"maxLength,omitempty"`
	MaxListSize json.Number  `json:"maxListSize,omitempty"`
	MaxValue    json.Number  `json:"maxValue,omitempty"`
	MinValue    json.Number  `json:"minValue,omitempty"`
	SubParams   []ApiParam   `json:"subParams,omitempty"`
	UsePointer  bool         `json:"use_pointer,omitempty"`
}

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
	case BYTES_PARAM_TYPE:
		param.Type = "*model.File"
	case JSON_PARAM_TYPE:
		param.Type = "string"
	default:
		paramType = strings.Title(paramType)
		if strings.HasSuffix(paramType, "Dto") {
			paramType = strings.TrimSuffix(paramType, "Dto") + "DTO"
		}
		if strings.HasSuffix(paramType, "Do") {
			paramType = strings.TrimSuffix(paramType, "Do") + "DO"
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
			for _, par := range p.SubParams {
				param.Obj = append(param.Obj, par.TplParam(apiName))
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

func convertConflictModelName(apiName string, name string) string {
	for _, n := range ConflictModels {
		if name == n {
			return fmt.Sprintf("%s%s", apiName, n)
		}
	}
	return name
}

func clearDesc(desc string, inline bool) string {
	desc = strings.ReplaceAll(desc, "/*", "")
	desc = strings.ReplaceAll(desc, "*/", "")
	if inline {
		desc = strings.ReplaceAll(desc, "\n", "")
	}
	return desc
}
