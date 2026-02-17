class YearProgress < Formula
  desc "Display year progress as a colorful ASCII progress bar"
  homepage "https://github.com/haroldelopez/year-progress-plugin"
  url "https://github.com/haroldelopez/year-progress-plugin/archive/refs/tags/v1.0.5.tar.gz"
  version "1.0.5"
  sha256 "50571b0d044f19862c7f16dd8eef2c4013d6e38f8ddb94873deb08c3e89ce028"
  license "MIT"
  head "https://github.com/haroldelopez/year-progress-plugin.git"

  depends_on "go" => :build

  def install
    system "go", "build", "-ldflags", "-s -w -X main.Version=#{version}", "-o", "year-progress", "."
    bin.install "year-progress"
  end

  test do
    output = shell_output("#{bin}/year-progress --version")
    assert_match "year-progress", output
  end
end
