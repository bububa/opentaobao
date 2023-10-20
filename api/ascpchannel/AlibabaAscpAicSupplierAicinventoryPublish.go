package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpaicsupplieraicinventorypublish 商家仓操作aic库存发布服务
// alibaba.ascp.aic.supplier.aicinventory.publish
//
// 商家调用这个接口来发布增加库存数据
func Alibabaascpaicsupplieraicinventorypublish(clt *core.SDKClient, req *ascpchannel.AlibabaascpaicsupplieraicinventorypublishAPIRequest, session string) (*ascpchannel.AlibabaascpaicsupplieraicinventorypublishAPIResponse, error) {
	var resp ascpchannel.AlibabaascpaicsupplieraicinventorypublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
