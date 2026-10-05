{{/*
Common labels applied to the Azure Service Operator DnsZonesSRVRecord resource.
*/}}
{{- define "azure-network-dnszonessrvrecord.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
