package network

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"
)

func eventsMapping(netsStatsList []NETstats) []common.MapStr {
	myEvents := []common.MapStr{}
	for _, netsStats := range netsStatsList {
		myEvents = append(myEvents, eventMapping(&netsStats))
	}
	return myEvents
}

func eventMapping(stats *NETstats) common.MapStr {
	event := common.MapStr{
		mb.ModuleData: common.MapStr{
			"container": stats.Container.ToMapStr(),
		},
		"interface": stats.NameInterface,
		"in": common.MapStr{
			"bytes":   stats.RxBytes,
			"dropped": stats.RxDropped,
			"errors":  stats.RxErrors,
			"packets": stats.RxPackets,
		},
		"out": common.MapStr{
			"bytes":   stats.TxBytes,
			"dropped": stats.TxDropped,
			"errors":  stats.TxErrors,
			"packets": stats.TxPackets,
		},
	}
	return event
}
