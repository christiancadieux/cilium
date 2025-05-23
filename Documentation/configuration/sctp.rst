.. only:: not (epub or latex or html)

    WARNING: You are looking at unreleased Cilium documentation.
    Please use the official rendered version released here:
    https://docs.cilium.io

.. _sctp:

*******************
SCTP support (beta)
*******************

.. include:: ../beta.rst

Enabling
========
Pass ``--set sctp.enabled=true`` to helm.

.. admonition:: Video
 :class: attention

  You can also watch a video explanation of Cilium's SCTP support in `eCHO episode 78: Stream Control Transmission Protocol (SCTP) <https://www.youtube.com/watch?v=2lD86qNHXXI>`__.

Limitations
===========
Cilium supports basic SCTP support. Specifically, the following is supported:
 - Pod <-> Pod communication
 - Pod <-> Service communication [*]
 - Pod <-> Pod communication with network policies applied to SCTP traffic [*]

.. note::
   [*] SCTP support does not support rewriting ports for SCTP packets. This means
   that when defining services, the targetPort **MUST** equal the port, otherwise
   the packet will be dropped.

.. warning::
   Cilium does not support the following for SCTP:
    - Multihoming
    - Policies for pod-to-VIP
    - :ref:`Kube-proxy replacement (KPR)<kubeproxy-free>` when port rewriting
      is necessary: for example, NodePort Services are not supported with the
      combination of KPR and SCTP.    
    - BPF masquerading
    - Egress gateway
