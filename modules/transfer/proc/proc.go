package proc

import (
	log "github.com/sirupsen/logrus"
	nproc "github.com/toolkits/proc"
)

// trace
var (
	RecvDataTrace = nproc.NewDataTrace("RecvDataTrace", 3)
)

// filter
var (
	RecvDataFilter = nproc.NewDataFilter("RecvDataFilter", 5)
)

// 统计指标的整体数据
var (
	// 计数统计,正确计数,错误计数, ...
	RecvCnt       = nproc.NewSCounterQps("RecvCnt")
	RpcRecvCnt    = nproc.NewSCounterQps("RpcRecvCnt")
	HttpRecvCnt   = nproc.NewSCounterQps("HttpRecvCnt")
	SocketRecvCnt = nproc.NewSCounterQps("SocketRecvCnt")

	SendToJudgeCnt      = nproc.NewSCounterQps("SendToJudgeCnt")
	SendToTsdbCnt       = nproc.NewSCounterQps("SendToTsdbCnt")
	SendToGraphCnt      = nproc.NewSCounterQps("SendToGraphCnt")
	SendToInfluxdbCnt   = nproc.NewSCounterQps("SendToInfluxdbCnt")
	SendToNqmIcmpCnt    = nproc.NewSCounterQps("SendToNqmIcmpCnt")
	SendToNqmTcpCnt     = nproc.NewSCounterQps("SendToNqmTcpCnt")
	SendToNqmTcpconnCnt = nproc.NewSCounterQps("SendToNqmTcpconnCnt")
	SendToStagingCnt    = nproc.NewSCounterQps("SendToStagingCnt")

	SendToJudgeDropCnt      = nproc.NewSCounterQps("SendToJudgeDropCnt")
	SendToTsdbDropCnt       = nproc.NewSCounterQps("SendToTsdbDropCnt")
	SendToGraphDropCnt      = nproc.NewSCounterQps("SendToGraphDropCnt")
	SendToInfluxdbDropCnt   = nproc.NewSCounterQps("SendToInfluxdbDropCnt")
	SendToNqmIcmpDropCnt    = nproc.NewSCounterQps("SendToNqmIcmpDropCnt")
	SendToNqmTcpDropCnt     = nproc.NewSCounterQps("SendToNqmTcpDropCnt")
	SendToNqmTcpconnDropCnt = nproc.NewSCounterQps("SendToNqmTcpconnDropCnt")
	SendToStagingDropCnt    = nproc.NewSCounterQps("SendToStagingDropCnt")

	SendToJudgeFailCnt      = nproc.NewSCounterQps("SendToJudgeFailCnt")
	SendToTsdbFailCnt       = nproc.NewSCounterQps("SendToTsdbFailCnt")
	SendToGraphFailCnt      = nproc.NewSCounterQps("SendToGraphFailCnt")
	SendToInfluxdbFailCnt   = nproc.NewSCounterQps("SendToInfluxdbFailCnt")
	SendToNqmIcmpFailCnt    = nproc.NewSCounterQps("SendToNqmIcmpFailCnt")
	SendToNqmTcpFailCnt     = nproc.NewSCounterQps("SendToNqmTcpFailCnt")
	SendToNqmTcpconnFailCnt = nproc.NewSCounterQps("SendToNqmTcpconnFailCnt")
	SendToStagingFailCnt    = nproc.NewSCounterQps("SendToStagingFailCnt")

	// 发送缓存大小
	JudgeQueuesCnt    = nproc.NewSCounterBase("JudgeSendCacheCnt")
	TsdbQueuesCnt     = nproc.NewSCounterBase("TsdbSendCacheCnt")
	GraphQueuesCnt    = nproc.NewSCounterBase("GraphSendCacheCnt")
	InfluxdbQueuesCnt = nproc.NewSCounterBase("InfluxdbSendCacheCnt")
	NqmRpcQueuesCnt   = nproc.NewSCounterBase("NqmRpcSendCacheCnt")
	StagingQueuesCnt  = nproc.NewSCounterBase("StagingSendCacheCnt")
)

func Start() {
	log.Println("proc.Start, ok")
}

func GetAll() []interface{} {
	ret := make([]interface{}, 0)

	// recv cnt
	ret = append(ret, RecvCnt.Get())
	ret = append(ret, RpcRecvCnt.Get())
	ret = append(ret, HttpRecvCnt.Get())
	ret = append(ret, SocketRecvCnt.Get())

	// send cnt
	ret = append(ret, SendToJudgeCnt.Get())
	ret = append(ret, SendToTsdbCnt.Get())
	ret = append(ret, SendToGraphCnt.Get())
	ret = append(ret, SendToInfluxdbCnt.Get())
	ret = append(ret, SendToNqmIcmpCnt.Get())
	ret = append(ret, SendToNqmTcpCnt.Get())
	ret = append(ret, SendToNqmTcpconnCnt.Get())
	ret = append(ret, SendToStagingCnt.Get())

	// drop cnt
	ret = append(ret, SendToJudgeDropCnt.Get())
	ret = append(ret, SendToTsdbDropCnt.Get())
	ret = append(ret, SendToGraphDropCnt.Get())
	ret = append(ret, SendToInfluxdbDropCnt.Get())
	ret = append(ret, SendToNqmIcmpDropCnt.Get())
	ret = append(ret, SendToNqmTcpDropCnt.Get())
	ret = append(ret, SendToNqmTcpconnDropCnt.Get())
	ret = append(ret, SendToStagingDropCnt.Get())

	// send fail cnt
	ret = append(ret, SendToJudgeFailCnt.Get())
	ret = append(ret, SendToTsdbFailCnt.Get())
	ret = append(ret, SendToGraphFailCnt.Get())
	ret = append(ret, SendToInfluxdbFailCnt.Get())
	ret = append(ret, SendToNqmIcmpFailCnt.Get())
	ret = append(ret, SendToNqmTcpFailCnt.Get())
	ret = append(ret, SendToNqmTcpconnFailCnt.Get())
	ret = append(ret, SendToStagingFailCnt.Get())

	// cache cnt
	ret = append(ret, JudgeQueuesCnt.Get())
	ret = append(ret, TsdbQueuesCnt.Get())
	ret = append(ret, GraphQueuesCnt.Get())
	ret = append(ret, InfluxdbQueuesCnt.Get())
	ret = append(ret, NqmRpcQueuesCnt.Get())
	ret = append(ret, StagingQueuesCnt.Get())

	return ret
}
