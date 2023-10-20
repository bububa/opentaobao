package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// Alibabawestcrmgradeget 获取等级列表
// alibaba.westcrm.grade.get
//
// 获取会员卡等级列表
func Alibabawestcrmgradeget(clt *core.SDKClient, req *westcrm.AlibabawestcrmgradegetAPIRequest, session string) (*westcrm.AlibabawestcrmgradegetAPIResponse, error) {
	var resp westcrm.AlibabawestcrmgradegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
