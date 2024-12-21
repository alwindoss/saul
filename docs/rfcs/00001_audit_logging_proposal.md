# Audit logging proposal

## Introduction

The Integrating the Healthcare Enterprise (IHE) Audit Trail and Node Authentication (ATNA) profile establishes security and auditing standards to ensure secure communication and data exchange in healthcare environments. This proposal outlines the development of an ATNA-compliant audit logger that adheres to IHE standards and provides robust security and audit capabilities for healthcare systems.

## Objective

- Develop a secure, scalable, and ATNA-compliant audit logging system.
- Ensure compatibility with healthcare applications and systems to achieve seamless integration.
- Provide a reliable mechanism for logging, transmitting, and reviewing audit events.

## Scope

The proposed audit logger will support:

1. IHE ATNA Profile Compliance: Implement all necessary features to conform to the IHE ATNA profile, including secure logging and transmission of events.
2. Audit Event Logging: Capture events such as user access, data modifications, and system interactions in compliance with DICOM and HL7 standards.
3. Secure Communication: Use transport-layer encryption (TLS) for transmitting audit logs to ensure data integrity and confidentiality.
4. Centralized Logging: Support centralized audit repositories for multi-system environments.
5. Integration Capabilities: Provide APIs and interfaces for integration with existing healthcare systems and applications.
6. Monitoring and Reporting: Offer dashboards and tools for real-time monitoring, querying, and reporting of audit logs.

## Need for Open Source

An ATNA-compliant audit logger that adheres to IHE standards represents a critical need in the Healthcare ecosystem. It is essential that the source code for the service itself is auditable.

## Key Features

1. Secure Event Logging

- Support structured logging formats like XML and JSON.
- Include metadata such as timestamps, user IDs, event types, and outcome status.

2. ATNA-Compliant Transport

- Use Syslog over TLS (RFC 5425) for secure audit log transmission.
- Implement mutual TLS for client-server authentication.

3. Configurable Policies

- Allow administrators to define what events are logged and how long logs are retained.

4. Scalability

- Support distributed healthcare environments with high availability and fault tolerance.

5. Standards Compliance

- Conform to HL7, DICOM, and IHE specifications for interoperability.
- Support event formats defined by the IHE IT Infrastructure Technical Framework (ITI TF).

6. Audit Viewer

- Provide tools for visualizing and analyzing audit logs.
- Enable secure access control to audit data for compliance purposes.

## Technical Approach

1. Architecture:

- A microservices-based architecture to ensure scalability and modularity.
Core components: Audit log collector, transport service, storage backend, and audit viewer.

2. Storage:

- Use a secure, high-performance database (e.g., PostgreSQL, Elasticsearch) for storing logs.
Ensure compliance with HIPAA and GDPR for data privacy.

3. Development Stack:

- Programming Language: Go for high performance and concurrency support.
- Libraries: OpenSSL or equivalent for encryption, Go-based logging frameworks for log management.

4. Security Measures:

- Encrypt audit logs at rest and in transit.
- Implement role-based access control (RBAC) and audit trail integrity checks.
