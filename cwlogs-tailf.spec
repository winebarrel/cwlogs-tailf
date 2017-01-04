%define  debug_package %{nil}

Name:   cwlogs-tailf
Version:  0.1.4
Release:  1%{?dist}
Summary:  Follow tail messages of CloudWatch Logs.

Group:    Development/Tools
License:  MIT
URL:    https://github.com/winebarrel/cwlogs-tailf
Source0:  %{name}.tar.gz
# https://github.com/winebarrel/cwlogs-tailf/releases/download/v%{version}/cwlogs-tailf_%{version}.tar.gz

%description
Follow tail messages of CloudWatch Logs.

%prep
%setup -q -n src

%build
make

%install
rm -rf %{buildroot}
mkdir -p %{buildroot}/usr/bin
install -m 755 cwlogs-tailf %{buildroot}/usr/bin/

%files
%defattr(755,root,root,-)
/usr/bin/cwlogs-tailf
