{{/*
Common labels applied to the Azure Service Operator ServersDatabasesAdvancedThreatProtectionSetting resource.
*/}}
{{- define "azure-sql-serversdatabasesadvancedthreatprotectionsetting.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
