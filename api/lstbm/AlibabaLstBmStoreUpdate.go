package lstbm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstbm"
)

/* AlibabaLstBmStoreUpdate
修改品牌商自有门店数据
alibaba.lst.bm.store.update

修改品牌商自有门店数据 */
func AlibabaLstBmStoreUpdate(clt *core.SDKClient, req *lstbm.AlibabaLstBmStoreUpdateAPIRequest, session string) (*lstbm.AlibabaLstBmStoreUpdateAPIResponse, error) {
	var resp lstbm.AlibabaLstBmStoreUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
