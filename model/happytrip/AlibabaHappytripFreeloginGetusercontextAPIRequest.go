package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripFreeloginGetusercontextAPIRequest
提供给外部系统的免登校验 API请求
alibaba.happytrip.freelogin.getusercontext

免登融合，提供免登相关接口给外部供应商做登录验证 */
type AlibabaHappytripFreeloginGetusercontextAPIRequest struct {
	model.Params
	// 请求入参
	_req *SsoParamDto
}

// New
