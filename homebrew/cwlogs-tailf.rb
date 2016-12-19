require 'formula'

class CwlogsTailf < Formula
  VERSION = '0.1.4'

  homepage 'https://github.com/winebarrel/cwlogs-tailf'
  url "https://github.com/winebarrel/cwlogs-tailf/releases/download/v#{VERSION}/cwlogs-tailf-v#{VERSION}-darwin-amd64.gz"
  sha256 '535c4e3d9accd74e5cc0994c4fd827330f42412d19cd090bd184e46d02aa8399'
  version VERSION
  head 'https://github.com/winebarrel/cwlogs-tailf.git', :branch => 'master'

  def install
    system "mv cwlogs-tailf-v#{VERSION}-darwin-amd64 cwlogs-tailf"
    bin.install 'cwlogs-tailf'
  end
end
