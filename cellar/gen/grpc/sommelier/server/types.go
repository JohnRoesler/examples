// Code generated by goa v2.2.6, DO NOT EDIT.
//
// sommelier gRPC server types
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package server

import (
	sommelierpb "goa.design/examples/cellar/gen/grpc/sommelier/pb"
	sommelier "goa.design/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
)

// NewPickPayload builds the payload of the "pick" endpoint of the "sommelier"
// service from the gRPC request type.
func NewPickPayload(message *sommelierpb.PickRequest) *sommelier.Criteria {
	v := &sommelier.Criteria{}
	if message.Name != "" {
		v.Name = &message.Name
	}
	if message.Winery != "" {
		v.Winery = &message.Winery
	}
	if message.Varietal != nil {
		v.Varietal = make([]string, len(message.Varietal))
		for i, val := range message.Varietal {
			v.Varietal[i] = val
		}
	}
	return v
}

// NewStoredBottleCollection builds the gRPC response type from the result of
// the "pick" endpoint of the "sommelier" service.
func NewStoredBottleCollection(result sommelierviews.StoredBottleCollectionView) *sommelierpb.StoredBottleCollection {
	message := &sommelierpb.StoredBottleCollection{}
	message.Field = make([]*sommelierpb.StoredBottle, len(result))
	for i, val := range result {
		message.Field[i] = &sommelierpb.StoredBottle{}
		if val.ID != nil {
			message.Field[i].Id = *val.ID
		}
		if val.Name != nil {
			message.Field[i].Name = *val.Name
		}
		if val.Vintage != nil {
			message.Field[i].Vintage = *val.Vintage
		}
		if val.Description != nil {
			message.Field[i].Description = *val.Description
		}
		if val.Rating != nil {
			message.Field[i].Rating = *val.Rating
		}
		if val.Winery != nil {
			message.Field[i].Winery = svcSommelierviewsWineryViewToSommelierpbWinery(val.Winery)
		}
		if val.Composition != nil {
			message.Field[i].Composition = make([]*sommelierpb.Component, len(val.Composition))
			for j, val := range val.Composition {
				message.Field[i].Composition[j] = &sommelierpb.Component{}
				if val.Varietal != nil {
					message.Field[i].Composition[j].Varietal = *val.Varietal
				}
				if val.Percentage != nil {
					message.Field[i].Composition[j].Percentage = *val.Percentage
				}
			}
		}
	}
	return message
}

// svcSommelierviewsWineryViewToSommelierpbWinery builds a value of type
// *sommelierpb.Winery from a value of type *sommelierviews.WineryView.
func svcSommelierviewsWineryViewToSommelierpbWinery(v *sommelierviews.WineryView) *sommelierpb.Winery {
	res := &sommelierpb.Winery{}
	if v.Name != nil {
		res.Name = *v.Name
	}
	if v.Region != nil {
		res.Region = *v.Region
	}
	if v.Country != nil {
		res.Country = *v.Country
	}
	if v.URL != nil {
		res.Url = *v.URL
	}

	return res
}

// protobufSommelierpbWineryToSommelierviewsWineryView builds a value of type
// *sommelierviews.WineryView from a value of type *sommelierpb.Winery.
func protobufSommelierpbWineryToSommelierviewsWineryView(v *sommelierpb.Winery) *sommelierviews.WineryView {
	res := &sommelierviews.WineryView{
		Name:    &v.Name,
		Region:  &v.Region,
		Country: &v.Country,
	}
	if v.Url != "" {
		res.URL = &v.Url
	}

	return res
}
