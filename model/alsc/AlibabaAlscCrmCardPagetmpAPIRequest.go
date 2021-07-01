package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardPagetmpAPIRequest
查询卡模板列表(支持数据下行) API请求
alibaba.alsc.crm.card.pagetmp

查询卡模板列表(支持数据下行)
当传递了数据下行参数:
     * isDeleted,lastMaxId,gmtModified,num时,进行数据下行处理,返回结果不带分页信息
     * 否则分页查询卡模板,返回结果带有分页信息 */
type AlibabaAlscCrmCardPagetmpAPIRequest struct {
	model.Params
	// 请求结果
	_paramPullCardTemplateOpenReq *PullCardTemplateOpenReq
}

// New
