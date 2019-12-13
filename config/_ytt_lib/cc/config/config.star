load("@ytt:data", "data")

def uaa_client(): return data.values.uaa.client

def url(): return "uaa." + data.values.namespace + ".svc.cluster.local"
