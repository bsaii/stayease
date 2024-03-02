data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/model",
    "--dialect", "postgres",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}