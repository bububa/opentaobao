package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* AlibabaAlscCrmCardPagetmp
查询卡模板列表(支持数据下行)
alibaba.alsc.crm.card.pagetmp

查询卡模板列表(支持数据下行)
当传递了数据下行参数:
     * isDeleted,lastMaxId,gmtModified,num时,进行数据下行处理,返回结果不带分页信息
     * 否则分页查询卡模板,返回结果带有分页信息 */
func AlibabaAlscCrmCardPagetmp(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardPagetmpAPIRequest, session string) (*alsc.AlibabaAlscCrmCardPagetmpAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmCardPagetmpAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
