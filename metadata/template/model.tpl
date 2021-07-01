package {{ .Pkg }}

{{- if eq .ImportModel true}}
import (
    "github.com/bububa/opentaobao/model"
)
{{- end }}

// {{ .Name }} 结构体
type {{ .Name }} struct {
{{- range $v := .Params }}
    // {{ $v.Desc }}
    {{- if and (eq $v.IsList true) }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty" xml:"{{ $v.ParamKey }}>{{ $v.SnakeType }},omitempty"`
    {{- else }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty" xml:"{{ $v.ParamKey }},omitempty"`
    {{- end }}
{{- end }}
}
