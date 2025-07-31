# 🔐 Security Policy

## 📅 Supported Versions

| Version            | Status     |
|--------------------|------------|
| 1.0.0 (Current)    | ✅ Supported |
| < 1.0.0            | ❌ Unsupported |

---

## 🚨 Reporting a Vulnerability

If you discover a security vulnerability in **QuantumHybridCompiler**, please **do not create a public issue**.

Instead, follow this secure reporting process:

- 📧 Email: **security@quantumhybrid.org**
- 🔒 GPG Key: [Public Key Placeholder]  
- 📁 You may encrypt your report and include reproduction steps or PoC if possible.

We aim to respond within **72 hours**, and resolve verified issues as quickly as possible.

---

## 🔍 Disclosure Policy

We follow a **coordinated vulnerability disclosure** process:

1. Report is received and acknowledged.
2. Maintainers validate the issue and assess impact.
3. A fix is prepared and tested.
4. Public disclosure is coordinated with the reporter (if applicable).
5. A security advisory is published via [GitHub Security Advisories](https://github.com/AUSP59/QuantumHybridCompiler/security/advisories).

---

## 🔐 Security Features

QuantumHybridCompiler implements:

- ✅ DSL input sanitization
- ✅ Isolation of quantum translation layer
- ✅ STRIDE-based threat modeling
- ✅ SPDX-compliant SBOM (`sbom/`)
- ✅ In-Toto attestation (`sbom/attestation.intoto.jsonl`)
- ✅ `govulncheck`, `gosec`, `golangci-lint` integration
- ✅ Reproducible builds (via `Makefile`, Docker)
- ✅ Secure supply chain practices (SLSA Level 1 ready)

---

## ✅ Best Practices for Users

- Always use **verified releases** from GitHub or Docker Hub.
- Avoid executing untrusted `.qdsl` code from unknown sources.
- Enable CI checks (`make lint`, `make test`) before deploying.
- Use isolated environments when testing external DSL programs.
- Monitor this repo for future security advisories.

---

## 🤝 Community Security

We value responsible disclosure and collaboration.  
If you’re a security researcher, we welcome your audit and expertise.

Let’s make open source safer — together.

— The QuantumHybridCompiler Team
