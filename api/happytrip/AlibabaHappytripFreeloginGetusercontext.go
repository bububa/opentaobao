package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

/* AlibabaHappytripFreeloginGetusercontext
提供给外部系统的免登校验
alibaba.happytrip.freelogin.getusercontext

免登融合，提供免登相关接口给外部供应商做登录验证 */
func AlibabaHappytripFreeloginGetusercontext(clt *core.SDKClient, req *happytrip.AlibabaHappytripFreeloginGetusercontextAPIRequest, session string) (*happytrip.AlibabaHappytripFreeloginGetusercontextAPIResponse, error) {
	var resp happytrip.AlibabaHappytripFreeloginGetusercontextAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
