package {{ .Pkg }}

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/{{ .Pkg }}"
)

/* 
{{ .ChineseName }} 
{{ .ApiName }}

{{ .Desc }}
*/
func {{ .Name }}(clt *core.SDKClient, req *{{ .Pkg }}.{{ .Name }}Request, session string) (*{{ .Pkg }}.{{ .Name }}Response, error) {
    var resp {{ .Pkg }}.{{ .Name }}APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
