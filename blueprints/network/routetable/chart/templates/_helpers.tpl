{{/*
Common labels applied to the Azure Service Operator RouteTable resource.
*/}}
{{- define "azure-network-routetable.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
