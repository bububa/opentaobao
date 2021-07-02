package baoxian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baoxian"
)

// AlipayBaoxianClaimUploadattachment 资料上传接口
// alipay.baoxian.claim.uploadattachment
//
// 给合作伙伴上传申请理赔材料
func AlipayBaoxianClaimUploadattachment(clt *core.SDKClient, req *baoxian.AlipayBaoxianClaimUploadattachmentAPIRequest, session string) (*baoxian.AlipayBaoxianClaimUploadattachmentAPIResponse, error) {
	var resp baoxian.AlipayBaoxianClaimUploadattachmentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
