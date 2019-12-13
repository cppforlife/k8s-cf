load("@ytt:data", "data")
load("@ytt:assert", "assert")

def value(*keys, values=None):
  val = values or data.values
  for k in keys:
  	val = getattr(val, k)
  end
  val and len(val) > 0 or assert.fail("missing/empty data value " + ".".join(keys))
  return val
end

value("uaa", "db", "password")

value("cc", "uaa_client_secret")
value("cc", "db", "password")

value("networking", "uaa_client_secret")
