// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	externalRef0Svr "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/server"
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	externalRef1Svr "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	externalRef2Svr "github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/server"
	externalRef2 "github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"net/http"
	"regexp"
)

var re = regexp.MustCompile(`[0-9a-z\-\._]+`)

const undefined = "undefined"

// GnmiPatchBody contains all the info required to build a gNMI requests
type GnmiPatchBody struct {
	Updates       []*gnmi.Update
	Deletes       []*gnmi.Path
	DefaultTarget string
	Ext100Name    *string
	Ext101Version *string
	Ext102Type    *string
	Ext110Info    *struct {
		ID    *string `json:"ID,omitempty"`
		Index *int    `json:"index,omitempty"`
	}
	Ext111Strategy *int
}

func encodeToGnmiPatchBody(jsonObj *types.PatchBody) (*GnmiPatchBody, error) {
	updates := make([]*gnmi.Update, 0)
	deletes := make([]*gnmi.Path, 0)

	pb := &GnmiPatchBody{}

	if jsonObj.Extensions != nil {
		pb.Ext100Name = jsonObj.Extensions.ChangeName100
		pb.Ext101Version = jsonObj.Extensions.ModelVersion101
		pb.Ext102Type = jsonObj.Extensions.ModelType102
		//pb.Ext110Info = jsonObj.Extensions.TransactionInfo110
		pb.Ext111Strategy = jsonObj.Extensions.TransactionStrategy111
	}

	if !re.MatchString(jsonObj.DefaultTarget) {
		return nil, fmt.Errorf("default-target cannot be blank")
	}
	pb.DefaultTarget = jsonObj.DefaultTarget

	gnmiUpdates, err := encodeToGnmiElements(jsonObj.Updates, jsonObj.DefaultTarget, false)
	if err != nil {
		return nil, err
	}
	updates = append(updates, gnmiUpdates...)
	pb.Updates = updates

	gnmiDeletes, err := encodeToGnmiElements(jsonObj.Deletes, jsonObj.DefaultTarget, true)
	if err != nil {
		return nil, err
	}
	for _, gd := range gnmiDeletes {
		deletes = append(deletes, gd.Path)
	}
	pb.Deletes = deletes

	return pb, nil
}

func encodeToGnmiElements(elements *types.Elements, target string, forDelete bool) ([]*gnmi.Update, error) {
	if elements == nil {
		return nil, nil
	}
	updates := make([]*gnmi.Update, 0)

	// Aether 2.0.x

	if elements.ConnectivityServices200 != nil {
		connectivityServiceUpdates, err := externalRef0Svr.EncodeToGnmiConnectivityServices(
			elements.ConnectivityServices200, false, forDelete, externalRef0.Target(target),
			"/connectivity-services")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiConnectivityService() %s", err)
		}
		updates = append(updates, connectivityServiceUpdates...)
	}

	if elements.Enterprises200 != nil {
		for _, e := range *elements.Enterprises200.Enterprise {

			if e.EnterpriseId == undefined {
				log.Warnw("EnterpriseId is undefined", "enterprise", e)
				return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "enterprise-id-cannot-be-undefined")
			}

			if e.Site != nil && len(*e.Site) > 0 {
				for _, s := range *e.Site {
					if s.SiteId == undefined {
						log.Warnw("SiteId is undefined", "site", s)
						return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "site-id-cannot-be-undefined")
					}
				}
			}
		}

		enterpriseUpdates, err := externalRef0Svr.EncodeToGnmiEnterprises(
			elements.Enterprises200, false, forDelete, externalRef0.Target(target),
			"/enterprises")

		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiEnterprise() %s", err)
		}
		updates = append(updates, enterpriseUpdates...)
	}

	// Aether 2.1.x

	if elements.Application210 != nil && len(*elements.Application210) > 0 {
		applicationUpdates, err := externalRef1Svr.EncodeToGnmiApplicationList(elements.Application210, false, false, "", "/application")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiApplicationList() %s", err)
		}
		updates = append(updates, applicationUpdates...)
	}

	if elements.Site210 != nil && len(*elements.Site210) > 0 {
		siteUpdates, err := externalRef1Svr.EncodeToGnmiSiteList(elements.Site210, false, false, "", "/site")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiSiteList() %s", err)
		}
		updates = append(updates, siteUpdates...)
	}

	if elements.Template210 != nil && len(*elements.Template210) > 0 {
		templateUpdates, err := externalRef1Svr.EncodeToGnmiTemplateList(elements.Template210, false, false, "", "/template")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiTemplateList() %s", err)
		}
		updates = append(updates, templateUpdates...)
	}

	if elements.TrafficClass210 != nil && len(*elements.TrafficClass210) > 0 {
		trafficClassUpdates, err := externalRef1Svr.EncodeToGnmiTrafficClassList(elements.TrafficClass210, false, false, "", "/traffic-class")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiTrafficClassList() %s", err)
		}
		updates = append(updates, trafficClassUpdates...)
	}

	// Aether 4.x

	if elements.ConnectivityService400 != nil {
		connectivityServiceUpdates, err := externalRef2Svr.EncodeToGnmiConnectivityService(
			elements.ConnectivityService400, false, forDelete, externalRef2.Target(target),
			"/connectivity-service")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiConnectivityService() %s", err)
		}
		updates = append(updates, connectivityServiceUpdates...)
	}

	if elements.Enterprise400 != nil {
		enterpriseUpdates, err := externalRef2Svr.EncodeToGnmiEnterprise(
			elements.Enterprise400, false, forDelete, externalRef2.Target(target),
			"/enterprise")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiEnterprise() %s", err)
		}
		updates = append(updates, enterpriseUpdates...)
	}

	if elements.Application400 != nil {
		applicationUpdates, err := externalRef2Svr.EncodeToGnmiApplication(
			elements.Application400, false, forDelete, externalRef2.Target(target),
			"/application")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiApplication() %s", err)
		}
		updates = append(updates, applicationUpdates...)
	}

	if elements.DeviceGroup400 != nil {
		deviceGroupUpdates, err := externalRef2Svr.EncodeToGnmiDeviceGroup(
			elements.DeviceGroup400, false, forDelete, externalRef2.Target(target),
			"/device-group")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiDeviceGroup() %s", err)
		}
		updates = append(updates, deviceGroupUpdates...)
	}

	if elements.IpDomain400 != nil {
		ipDomainUpdates, err := externalRef2Svr.EncodeToGnmiIpDomain(
			elements.IpDomain400, false, forDelete, externalRef2.Target(target),
			"/ip-domain")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiIpDomain() %s", err)
		}
		updates = append(updates, ipDomainUpdates...)
	}

	if elements.Site400 != nil {
		siteUpdates, err := externalRef2Svr.EncodeToGnmiSite(
			elements.Site400, false, forDelete, externalRef2.Target(target),
			"/site")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiSite() %s", err)
		}
		updates = append(updates, siteUpdates...)
	}

	if elements.Template400 != nil {
		templateUpdates, err := externalRef2Svr.EncodeToGnmiTemplate(
			elements.Template400, false, forDelete, externalRef2.Target(target),
			"/template")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiTemplate() %s", err)
		}
		updates = append(updates, templateUpdates...)
	}

	if elements.TrafficClass400 != nil {
		trafficClassUpdates, err := externalRef2Svr.EncodeToGnmiTrafficClass(
			elements.TrafficClass400, false, forDelete, externalRef2.Target(target),
			"/traffic-class")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiTrafficClass() %s", err)
		}
		updates = append(updates, trafficClassUpdates...)
	}

	if elements.Upf400 != nil {
		upfUpdates, err := externalRef2Svr.EncodeToGnmiUpf(
			elements.Upf400, false, forDelete, externalRef2.Target(target),
			"/upf")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiUpf() %s", err)
		}
		updates = append(updates, upfUpdates...)
	}

	if elements.Vcs400 != nil {
		vcsUpdates, err := externalRef2Svr.EncodeToGnmiVcs(
			elements.Vcs400, false, forDelete, externalRef2.Target(target),
			"/vcs")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiVcs() %s", err)
		}
		updates = append(updates, vcsUpdates...)
	}

	return updates, nil
}
