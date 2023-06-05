# Create a temporary directory named "foo"
data "temporary_directory" "foo" {
  name = "foo"
}

# Use the temporary directory as output_path
data "archive_file" "foo" {
  type        = "zip"
  source_file = "${path.module}/init.tpl"
  output_path = data.temporary_directory.foo.id
}
