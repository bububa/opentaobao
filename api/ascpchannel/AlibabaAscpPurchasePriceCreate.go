package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascppurchasepricecreate ascp采购价写入接口
// alibaba.ascp.purchase.price.create
//
// 供应链平台采购价创建或修改接口
func Alibabaascppurchasepricecreate(clt *core.SDKClient, req *ascpchannel.AlibabaascppurchasepricecreateAPIRequest, session string) (*ascpchannel.AlibabaascppurchasepricecreateAPIResponse, error) {
	var resp ascpchannel.AlibabaascppurchasepricecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
