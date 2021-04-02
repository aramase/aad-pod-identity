package client

import (
	"encoding/json"
	"flag"
	"io"

	winio "github.com/Microsoft/go-winio"
	msg "github.com/Microsoft/hcnproxy/pkg/types"
	"github.com/sirupsen/logrus"
)

var (
	pipeName = flag.String("p", msg.PipeName, "path and name of the windows pipe")
	sddl     = flag.String("s", msg.SSDL, "security descriptor of the pipe")
)

// PipeDialerFunc is for mocking
type PipeDialerFunc func() (io.ReadWriteCloser, error)

// dialFunction is for depedency injection
var dialFunction = dialFunctionMain

// InvokeHNSRequest process the HNSRequest and returns the appropriate response.
func InvokeHNSRequest(req msg.HNSRequest) *msg.HNSResponse {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.InfoLevel)

	return processRequest(logrus.New(), &req, dialFunction)
}

func dialFunctionMain() (io.ReadWriteCloser, error) {
	conn, err := winio.DialPipe(*pipeName, nil)
	return conn, err
}

func processRequest(log *logrus.Logger, req *msg.HNSRequest, dialerFunc PipeDialerFunc) *msg.HNSResponse {
	resp := &msg.HNSResponse{}
	c, err := dialerFunc()
	if err != nil {
		log.WithError(err).Warn("Couldn't connect to the new pipe")
		return errorResponse(resp, err)
	}
	defer c.Close()

	w := json.NewEncoder(c)
	if err = w.Encode(req); err != nil {
		log.WithError(err).Warn("Could not write HNSRequest")
		return errorResponse(resp, err)
	}

	r := json.NewDecoder(c)
	if err = r.Decode(&resp); err != nil {
		log.WithError(err).Warn("Could not read HNSResponse")
		return errorResponse(resp, err)
	}
	return resp
}

func errorResponse(resp *msg.HNSResponse, err error) *msg.HNSResponse {
	resp.Error = err
	return resp
}
