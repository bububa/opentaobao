package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeRegionInfoSubmit 商圈专家信息同步
// alibaba.alihouse.existinghome.region.info.submit
//
// 商圈专家信息同步
func AlibabaAlihouseExistinghomeRegionInfoSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
