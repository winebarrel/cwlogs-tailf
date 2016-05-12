require 'formula'

class CwlogsTailf < Formula
  VERSION = '0.1.0'

  homepage 'https://github.com/winebarrel/cwlogs-tailf'
  url "https://github.com/winebarrel/cwlogs-tailf/releases/download/v#{VERSION}/cwlogs-tailf-v#{VERSION}-darwin-amd64.gz"
  sha256 'e81020b0c8059359d9b36dc9b967c41bca2f0b1a62139fe0d4979a833a2179af'
  version VERSION
  head 'https://github.com/winebarrel/cwlogs-tailf.git', :branch => 'master'

  def install
    system "mv cwlogs-tailf-v#{VERSION}-darwin-amd64 cwlogs-tailf"
    bin.install 'cwlogs-tailf'
  end
end
