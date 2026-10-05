{{/*
Common labels applied to the Azure Service Operator NetworkWatchersFlowLog resource.
*/}}
{{- define "azure-network-watchersflowlog.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
