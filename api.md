# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go">roarkanalytics</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go#HealthGetResponse">HealthGetResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go#HealthService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go">roarkanalytics</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go#HealthGetResponse">HealthGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CallAnalysis

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go">roarkanalytics</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go#CallAnalysisNewResponse">CallAnalysisNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go">roarkanalytics</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go#CallAnalysisGetResponse">CallAnalysisGetResponse</a>

Methods:

- <code title="post /v1/call-analysis">client.CallAnalysis.<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go#CallAnalysisService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go">roarkanalytics</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go#CallAnalysisNewParams">CallAnalysisNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go">roarkanalytics</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go#CallAnalysisNewResponse">CallAnalysisNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/call-analysis/{jobId}">client.CallAnalysis.<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go#CallAnalysisService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go">roarkanalytics</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/roark-analytics-go#CallAnalysisGetResponse">CallAnalysisGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
