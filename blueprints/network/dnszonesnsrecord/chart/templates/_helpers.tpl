{{/*
Common labels applied to the Azure Service Operator DnsZonesNSRecord resource.
*/}}
{{- define "azure-network-dnszonesnsrecord.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
