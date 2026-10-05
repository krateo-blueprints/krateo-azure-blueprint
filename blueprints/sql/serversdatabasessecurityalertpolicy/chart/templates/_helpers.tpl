{{/*
Common labels applied to the Azure Service Operator ServersDatabasesSecurityAlertPolicy resource.
*/}}
{{- define "azure-sql-serversdatabasessecurityalertpolicy.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
