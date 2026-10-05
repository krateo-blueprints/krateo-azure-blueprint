{{/*
Common labels applied to the Azure Service Operator ServersIPV6FirewallRule resource.
*/}}
{{- define "azure-sql-serversipv6firewallrule.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
