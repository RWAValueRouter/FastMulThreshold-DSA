package discv5

import "github.com/RWAValueRouter/FastMulThreshold-DSA/p2p/metrics"

var (
	ingressTrafficMeter = metrics.NewRegisteredMeter("discv5/InboundTraffic", nil)
	egressTrafficMeter  = metrics.NewRegisteredMeter("discv5/OutboundTraffic", nil)
)
