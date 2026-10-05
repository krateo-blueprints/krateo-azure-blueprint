{{/*
Common labels applied to the Azure Service Operator PublicIPAddress resource.
*/}}
{{- define "azure-network-publicipaddress.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
