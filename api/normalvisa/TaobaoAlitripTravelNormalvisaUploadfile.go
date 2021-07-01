package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

/* TaobaoAlitripTravelNormalvisaUploadfile
上传电子签证
taobao.alitrip.travel.normalvisa.uploadfile

上传电子签证 */
func TaobaoAlitripTravelNormalvisaUploadfile(clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaUploadfileAPIRequest, session string) (*normalvisa.TaobaoAlitripTravelNormalvisaUploadfileAPIResponse, error) {
	var resp normalvisa.TaobaoAlitripTravelNormalvisaUploadfileAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
