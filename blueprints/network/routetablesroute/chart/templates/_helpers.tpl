{{/*
Common labels applied to the Azure Service Operator RouteTablesRoute resource.
*/}}
{{- define "azure-network-routetablesroute.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
