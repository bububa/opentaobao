package {{ .Pkg }}

import (
    "sync"
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

// {{ .Name }}APIResponse {{ .ChineseName }} API返回值 
// {{ .ApiName }}
//
{{ html .Desc }}
type {{ .Name }}APIResponse struct {
    model.CommonResponse
    {{ .Name }}APIResponseModel
}

// Reset 清空结构体
func (m *{{ .Name }}APIResponse) Reset() {
    (&m.CommonResponse).Reset()
    (&m.{{ .Name }}APIResponseModel).Reset()
}

// {{ .Name }}APIResponseModel is {{ .ChineseName }} 成功返回结果
type {{ .Name }}APIResponseModel struct {
    XMLName xml.Name `xml:"{{ .ResponseKey }}"`
    {{- if eq .HasRequestId false }}
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         
    {{- end }}
{{- range $v := .ResponseParams }}
    // {{ html $v.Desc }}
    {{- if and (eq $v.IsList true) }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty" xml:"{{ $v.ParamKey }}>{{ $v.SnakeType }},omitempty"` 
    {{- else }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty" xml:"{{ $v.ParamKey }},omitempty"` 
    {{- end }}
{{- end }}
}

// Reset 清空结构体
func (m *{{ .Name }}APIResponseModel) Reset() {
    m.RequestId = ""
{{- range $v := .ResponseParams }}
    {{- if and (eq $v.IsList true) }}
      m.{{ $v.Name }} = m.{{ $v.Name }}[:0] 
    {{- else if and (eq $v.IsObject true) }}
      m.{{ $v.Name }} = nil
    {{- else if and (eq $v.IsNumber true) }}
      m.{{ $v.Name }} = 0
    {{- else if and (eq $v.IsBool true) }}
      m.{{ $v.Name }} = false
    {{- else }}
      m.{{ $v.Name }} = ""
    {{- end }}
{{- end }}
}

var pool{{ .Name }}APIResponse = sync.Pool{
    New: func() any {
      return new({{ .Name }}APIResponse)
    },
}

// Get{{ .Name }}APIResponse 从 sync.Pool 获取 {{ .Name }}APIResponse
func Get{{ .Name }}APIResponse() *{{ .Name }}APIResponse {
    return pool{{ .Name }}APIResponse.Get().(*{{ .Name }}APIResponse)
}

// Release{{ .Name }}APIResponse 将 {{ .Name }}APIResponse 保存到 sync.Pool
func Release{{ .Name }}APIResponse(v *{{ .Name }}APIResponse) {
    v.Reset()
    pool{{ .Name }}APIResponse.Put(v)
}

