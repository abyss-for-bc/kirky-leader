package report

import (
    `context`
)

type EventReportServer struct {

}

func (server *EventReportServer) Report(context.Context, *ReportRequest) (*ReportResponse, error){
    return nil, nil
}
func (server *EventReportServer)Poll(context.Context, *PollEvent) (*PollEvent, error){
    return nil, nil
}
