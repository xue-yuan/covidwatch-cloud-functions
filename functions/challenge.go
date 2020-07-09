package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"upload-token.functions/internal/pow"
	"upload-token.functions/internal/util"
)

// ChallengeHandler is a handler for the /challenge endpoint.
var ChallengeHandler = util.MakeHTTPHandler(challengeHandler)

// SubmitReportHandler is a handler for the /submitReport endpoint.
var SubmitReportHandler = util.MakeHTTPHandler(submitReportHandler)

func challengeHandler(ctx *util.Context) util.StatusError {
	if err := util.ValidateRequestMethod(ctx, "GET", ""); err != nil {
		return err
	}

	c, err := pow.GenerateChallenge(ctx)
	if err != nil {
		return util.NewInternalServerError(err)
	}
	json.NewEncoder(ctx.HTTPResponseWriter()).Encode(c)

	return nil
}

func submitReportHandler(ctx *util.Context) util.StatusError {
	if err := util.ValidateRequestMethod(ctx, "POST", ""); err != nil {
		return err
	}
	defer ctx.HTTPRequest().Body.Close()
	b, _ := ioutil.ReadAll(ctx.HTTPRequest().Body)

	var report util.Report
	err := json.NewDecoder(strings.NewReader(string(b))).Decode(&report)
	if err == nil {
		fmt.Println(report)
	}

	pow.StoreReport(report, ctx)
	return nil
}
