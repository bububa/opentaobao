package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaoglobalhandoverpdfget 获取面单PDF文件数据
// cainiao.global.handover.pdf.get
//
// 返回指定大包面单的PDF文件数据
func Cainiaoglobalhandoverpdfget(clt *core.SDKClient, req *cainiaohandover.CainiaoglobalhandoverpdfgetAPIRequest, session string) (*cainiaohandover.CainiaoglobalhandoverpdfgetAPIResponse, error) {
	var resp cainiaohandover.CainiaoglobalhandoverpdfgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
