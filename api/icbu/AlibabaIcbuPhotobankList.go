package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuphotobanklist 国际站图片银行查询接口
// alibaba.icbu.photobank.list
//
// 国际站图片银行查询接口
func Alibabaicbuphotobanklist(clt *core.SDKClient, req *icbu.AlibabaicbuphotobanklistAPIRequest, session string) (*icbu.AlibabaicbuphotobanklistAPIResponse, error) {
	var resp icbu.AlibabaicbuphotobanklistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
