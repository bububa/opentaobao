package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// Taobaoalitriptravelnormalvisauploadfile 上传电子签证
// taobao.alitrip.travel.normalvisa.uploadfile
//
// 上传电子签证
func Taobaoalitriptravelnormalvisauploadfile(clt *core.SDKClient, req *normalvisa.TaobaoalitriptravelnormalvisauploadfileAPIRequest, session string) (*normalvisa.TaobaoalitriptravelnormalvisauploadfileAPIResponse, error) {
	var resp normalvisa.TaobaoalitriptravelnormalvisauploadfileAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
