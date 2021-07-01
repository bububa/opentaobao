package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuPhotobankGroupOperateAPIRequest
图片银行分组操作接口 API请求
alibaba.icbu.photobank.group.operate

修改用户图片银行的分组信息，包括 新增分组，删除分组，分组重命名 */
type AlibabaIcbuPhotobankGroupOperateAPIRequest struct {
	model.Params
	// 图片分组操作请求对象
	_photoGroupOperationRequest *PhotoGroupOperationRequest
}

// New
