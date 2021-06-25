package {{ .Pkg }}

// {{ .Name }} 
type {{ .Name }} struct {
{{ range $v := .Params }}
    // {{ $v.Desc }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty"`
{{ end }}
}
