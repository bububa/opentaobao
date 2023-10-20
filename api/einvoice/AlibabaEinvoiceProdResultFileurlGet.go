package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceprodresultfileurlget 发票中台-发票文件下载地址查询
// alibaba.einvoice.prod.result.fileurl.get
//
// 发票文件下载地址查询，外部ISV通过该接口可以查对应发票文件
func Alibabaeinvoiceprodresultfileurlget(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceprodresultfileurlgetAPIRequest, session string) (*einvoice.AlibabaeinvoiceprodresultfileurlgetAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceprodresultfileurlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
