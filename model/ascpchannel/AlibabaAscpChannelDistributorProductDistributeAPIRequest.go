package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorProductDistributeAPIRequest 分销商基于渠道产品铺货到商品 API请求
// alibaba.ascp.channel.distributor.product.distribute
//
// 分销商基于渠道产品铺货到商品
type AlibabaAscpChannelDistributorProductDistributeAPIRequest struct {
	model.Params
	// 请求参数
	_itemDistributePublishRequest *ItemDistributePublishRequest
}

// NewAlibabaAscpChannelDistributorProductDistributeRequest 初始化AlibabaAscpChannelDistributorProductDistributeAPIRequest对象
func NewAlibabaAscpChannelDistributorProductDistributeRequest() *AlibabaAscpChannelDistributorProductDistributeAPIRequest {
	return &AlibabaAscpChannelDistributorProductDistributeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorProductDistributeAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.product.distribute"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorProductDistributeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelDistributorProductDistributeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemDistributePublishRequest is ItemDistributePublishRequest Setter
// 请求参数
func (r *AlibabaAscpChannelDistributorProductDistributeAPIRequest) SetItemDistributePublishRequest(_itemDistributePublishRequest *ItemDistributePublishRequest) error {
	r._itemDistributePublishRequest = _itemDistributePublishRequest
	r.Set("item_distribute_publish_request", _itemDistributePublishRequest)
	return nil
}

// GetItemDistributePublishRequest ItemDistributePublishRequest Getter
func (r AlibabaAscpChannelDistributorProductDistributeAPIRequest) GetItemDistributePublishRequest() *ItemDistributePublishRequest {
	return r._itemDistributePublishRequest
}
