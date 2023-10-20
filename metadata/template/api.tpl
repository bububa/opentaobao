package {{ .Pkg }}

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/{{ .Pkg }}"
)

// {{ .Name }} {{ .ChineseName }} 
// {{ .ApiName }}
//
{{ html .Desc }}
func {{ .Name }}(clt *core.SDKClient, req *{{ .Pkg }}.{{ .Name }}APIRequest, resp *{{ .Pkg }}.{{ .Name }}APIResponse, session string) error {
    return clt.Post(req, resp, session)
}
