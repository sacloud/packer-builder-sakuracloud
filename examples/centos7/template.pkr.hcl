source "sakuracloud" "example" {
  zone = "is1b"

  os_type   = "centos7"
  password  = "TestUserPassword01"
  disk_size = 20
  disk_plan = "ssd"

  core        = 2
  memory_size = 4

  archive_name        = "packer-example-centos"
  archive_description = "description of archive"
}

build {
  sources = [
    "source.sakuracloud.example"
  ]
  provisioner "shell" {
    inline = [
      "yum update -y",
      "curl -fsSL https://get.docker.com/ | sh",
      "systemctl enable docker.service",
    ]
  }
}

