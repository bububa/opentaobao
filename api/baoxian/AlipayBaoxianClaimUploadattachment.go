package baoxian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baoxian"
)

// Alipaybaoxianclaimuploadattachment 资料上传接口
// alipay.baoxian.claim.uploadattachment
//
// 给合作伙伴上传申请理赔材料
func Alipaybaoxianclaimuploadattachment(clt *core.SDKClient, req *baoxian.AlipaybaoxianclaimuploadattachmentAPIRequest, session string) (*baoxian.AlipaybaoxianclaimuploadattachmentAPIResponse, error) {
	var resp baoxian.AlipaybaoxianclaimuploadattachmentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
