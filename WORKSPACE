workspace(name = "amprobe")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository", "new_git_repository")

git_repository(
    name = "bazel_skylib",
    remote = "https://github.com/bazelbuild/bazel-skylib.git",
    tag = "0.7.0",  # change this to use a different release
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.18.6",
)

git_repository(
    name = "bazel_gazelle",
    commit = "a8a732e9f358d772df3884c0adeebba60274e780",
    remote = "https://github.com/bazelbuild/bazel-gazelle.git",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("@bazel_gazelle//:def.bzl", "go_repository")

go_repository(
    name = "com_github_prometheus_alertmanager",
    importpath = "github.com/prometheus/alertmanager",
    tag = "v0.17.0",
)

load("//:deps.bzl", "deps")

deps()
