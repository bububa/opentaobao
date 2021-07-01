package lstbm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstbm"
)

/* AlibabaLstBmStoreAdd
导入品牌商自有门店
alibaba.lst.bm.store.add

导入品牌商自有门店 */
func AlibabaLstBmStoreAdd(clt *core.SDKClient, req *lstbm.AlibabaLstBmStoreAddAPIRequest, session string) (*lstbm.AlibabaLstBmStoreAddAPIResponse, error) {
	var resp lstbm.AlibabaLstBmStoreAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
